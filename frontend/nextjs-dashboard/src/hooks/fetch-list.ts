import axios from 'axios'

const useCrud = () => {
  const post = async (route: string, values: any, token?: string) => {
    const response = await axios.post(
      `http://localhost:8080/${route}`,
      values,
      {
        headers: {
          Authorization: `Barear ${token}`,
        },
      }
    )

    if (response.status >= 200 && response.status < 300) {
      return response.data
    } else {
      throw Error('error')
    }
  }

  const get = async (route: string, token?: string) => {
    const response = await axios.get(`http://localhost:8080/${route}`, {
      headers: {
        Authorization: `Barear ${token}`,
      },
    })

    if (response.status >= 200 && response.status < 300) {
      return response.data
    } else {
      throw Error('error')
    }
  }

  const update = async (route: string, values?: any, token?: string) => {
    const response = await axios.patch(
      `http://localhost:8080/${route}`,
      values,
      {
        headers: {
          Authorization: `Barear ${token}`,
        },
      }
    )

    if (response.status >= 200 && response.status < 300) {
      return response.data
    } else {
      throw Error('error')
    }
  }

  const remove = async (route: string, token?: string) => {
    console.log(token)
    const response = await axios.delete(`http://localhost:8080/${route}`, {
      headers: {
        Authorization: `Barear ${token}`,
      },
    })

    if (response.status >= 200 && response.status < 300) {
      return response.data
    } else {
      throw Error('error')
    }
  }

  return { post, get, update, remove }
}

export default useCrud
