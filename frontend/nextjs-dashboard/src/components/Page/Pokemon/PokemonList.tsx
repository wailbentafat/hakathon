import {
  Dropdown, DropdownItem, DropdownMenu, DropdownToggle, Table,
} from 'react-bootstrap'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faEllipsisVertical } from '@fortawesome/free-solid-svg-icons'
import React from 'react'
import Image from 'next/image'
import Link from 'next/link'
import { Pokemon } from '@/models/pokemon'
import THSort from '@/components/TableSort/THSort'
import PokemonTypeLabel from '@/components/Page/Pokemon/PokemonTypeLabel'
import useDictionary from '@/locales/dictionary-hook'

type Props = {
  pokemons: Pokemon[];
}

export default function PokemonList(props: Props) {
  const { pokemons } = props
  const dict = useDictionary()

  return (
    <Table responsive bordered hover>
      <thead>
        <tr className="table-light dark:table-dark">
          <th>
            <THSort name="id">id</THSort>
          </th>
          <th>
            <THSort name="name">{dict.pokemons.attribute.name}</THSort>
          </th>
          <th>{dict.pokemons.attribute.type}</th>
          <th className="text-center">{dict.pokemons.attribute.egg_group}</th>
          <th className="text-end">
            <THSort name="hp">{dict.pokemons.attribute.hp}</THSort>
          </th>
          <th className="text-end">
            <THSort name="attack">{dict.pokemons.attribute.attack}</THSort>
          </th>
          <th className="text-end">
            <THSort name="special_defense">
              {dict.pokemons.attribute.speed}
            </THSort>
          </th>
          <th className="text-end">
            <THSort name="defense">{dict.pokemons.attribute.defense}</THSort>
          </th>
          <th className="text-end">
            <THSort name="special_attack">
              {dict.pokemons.attribute.sp_attack}
            </THSort>
          </th>
          <th className="text-end">
            <THSort name="special_defense">
              {dict.pokemons.attribute.sp_defense}
            </THSort>
          </th>
          <th aria-label="Action" />
        </tr>
      </thead>
      <tbody>
        {pokemons.map((pokemon) => (
          <tr key={pokemon.id}>
            <td>{pokemon.id}</td>

            <td>{pokemon.name}</td>
            <td>
              {pokemon.types.map((type) => (
                <span key={type.id} className="me-2">
                  <PokemonTypeLabel type={type} />
                </span>
              ))}
            </td>
            <td className="text-center" style={{ whiteSpace: 'pre' }}>
              {pokemon.egg_groups.map((eggGroup) => eggGroup.name).join('\n')}
            </td>
            <td className="text-end">{pokemon.hp}</td>
            <td className="text-end">{pokemon.attack}</td>
            <td className="text-end">{pokemon.defense}</td>
            <td className="text-end">{pokemon.special_attack}</td>
            <td className="text-end">{pokemon.special_defense}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  )
}
