import React, { useEffect, useState } from 'react'
import { Table } from 'react-bootstrap'
import THSort from '@/components/TableSort/THSort'
import useCrud from '@/hooks/fetch-list'


export default function PokemonList() {

  const {get} = useCrud()
  const token = ''
  const [data, setData] = useState([])

  useEffect(() => {
    const fetchingData = async () => {
      try {
        const res = await get('', token)
        console.log(res)
        setData(res)
      } catch (error) {
        console.log(error)
      }
    }

    //fetchingData()
  },[])


  // Sample fake data
  const fakeData = [
    {
      id: 1,
      name: 'John Doe',
      phone: '123-456-7890',
      bankCard: '**** **** **** 1234',
      location: 'New York, NY',
      category: 'Food',
      description: 'Regular customer, ordered pizza.',
      satisfied: 'Yes',
      staffId: 'STAFF001',
      date: '2024-09-20',
    },
    {
      id: 2,
      name: 'Jane Smith',
      phone: '987-654-3210',
      bankCard: '**** **** **** 5678',
      location: 'Los Angeles, CA',
      category: 'Clothing',
      description: 'Purchased summer collection.',
      satisfied: 'No',
      staffId: 'STAFF002',
      date: '2024-09-18',
    },
    {
      id: 3,
      name: 'Alice Johnson',
      phone: '555-555-5555',
      bankCard: '**** **** **** 0000',
      location: 'Chicago, IL',
      category: 'Electronics',
      description: 'Bought a new smartphone.',
      satisfied: 'Yes',
      staffId: 'STAFF003',
      date: '2024-09-15',
    },
  ]

  return (
    <Table responsive bordered hover>
      <thead>
        <tr className="table-light dark:table-dark">
          <th>
            <THSort name="id">id</THSort>
          </th>
          <th>
            <THSort name="name">Name</THSort>
          </th>
          <th>Phone number</th>
          <th className="text-center">Bank card</th>
          <th className="text-end">Location</th>
          <th className="text-end">
            <THSort name="category">Category</THSort>
          </th>
          <th className="text-end">Description</th>
          <th className="text-end">
            <THSort name="satisfied">Satisfied</THSort>
          </th>
          <th className="text-end">
            <THSort name="staff_id">Staff ID</THSort>
          </th>
          <th className="text-end">
            <THSort name="date">Date</THSort>
          </th>
          <th aria-label="Action" />
        </tr>
      </thead>
      <tbody>
        {fakeData.map((item) => (
          <tr key={item.id}>
            <td>{item.id}</td>
            <td>{item.name}</td>
            <td>{item.phone}</td>
            <td className="text-center">{item.bankCard}</td>
            <td className="text-end">{item.location}</td>
            <td className="text-end">{item.category}</td>
            <td className="text-end">{item.description}</td>
            <td className="text-end">{item.satisfied}</td>
            <td className="text-end">{item.staffId}</td>
            <td className="text-end">{item.date}</td>
            <td aria-label="Action">
              {/* You can add action buttons or a dropdown here */}
            </td>
          </tr>
        ))}
      </tbody>
    </Table>
  )
}
