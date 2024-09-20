import axios from 'axios';
import { NextAuthOptions, User } from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';
import { getDictionary } from '@/locales/dictionary';



export const authOptions: NextAuthOptions = {
  callbacks: {
    async jwt({ user, token }) {
      if (user) {
        // If a user object is returned, save the token to the token object
        token.accessToken = user.accessToken;
        return { ...token, user: { ...user as User } };
      }
      return token;
    },
   
  providers: [
    CredentialsProvider({
      credentials: {
        username: { type: 'text' },
        password: { type: 'password' },
      },
      async authorize(credentials) {
        if (!credentials) {
          return null;
        }
        const { username, password } = credentials;

        try {
          const response = await axios.post('http://your-backend-url/api/login', {
             email :username,
            password,
          });

          // Assuming the backend responds with a structure like:
          // { success: true, token: 'your_token', user: { ... } }
          if (response.data.success) {
            const token = response.data.token; // The token from your response

            return {
              accessToken: token, // Include the token for further use
            };
          } else {
            const dict = await getDictionary();
            throw new Error(dict.login.message.auth_failed);
          }
        } catch (error) {
          const dict = await getDictionary();
          throw new Error(dict.login.message.auth_failed || error.message);
        }
      },
    }),
  ],
};
