"use client"

import { useEffect, useState } from 'react'
import { Table } from 'react-bootstrap'
import THSort from '@/components/TableSort/THSort'
import axios from 'axios'

import Cookies from 'js-cookie'
function PokemonList() {

  const token = Cookies.get('token')
  const [data, setData] = useState([])
  
  useEffect(() => {
    const fetchingData =  () => {
      try {
        axios.get("http://localhost:8080/get_complain",{
          headers:{Authorization:`Bearer ${token} `}
        }).then((res)=>{
          console.log(res)
        }).catch((err)=>{
          console.log(err)
        })
      } catch (error) {
        console.log(error)
      }
    }
    fetchingData()
    
  },[])


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
        {/* {data.map((item) => (
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
            </td>
          </tr>
        ))} */}
      </tbody>
    </Table>
  )
}


export default PokemonList