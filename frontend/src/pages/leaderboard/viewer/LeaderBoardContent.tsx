import { Field, LeaderboardFull } from '@/types/leaderboard'
import { fieldToDisplayValue } from '@/utils/fieldToDisplayValue'
import { formatDuration } from '@/utils/formatDuration'
import { isDigitField } from '@/utils/isDigitField'
import React from 'react'

interface PropType {
  data: LeaderboardFull
}

const LeaderBoardContent: React.FC<PropType> = ({ data }) => {
  return (
    <div className="px-6 pb-6">
      <div className="overflow-x-auto">
        <table className="w-full text-sm text-left">
          <TableHeader fields={data.fields} />
          <tbody>
            {data.data.map((row, index) => (
              <TableRow row={row} index={index} fields={data.fields} />
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}


interface TableHeaderPropType {
  fields: Field[]
}

const TableHeader: React.FC<TableHeaderPropType> = ({ fields }) => {
  return (
    <thead className="text-xs text-gray-700 uppercase bg-white">
      <tr>
        <th scope="col" className="px-6 py-3 text-center w-12">#</th>
        {fields.map((field) => (
          <th key={field.fieldName} scope="col" className="px-6 py-3">
            {field.name}
          </th>
        ))}
      </tr>
    </thead>
  )
}


interface TableRowPropType {
  row: any
  index: number
  fields: Field[]
}

const TableRow: React.FC<TableRowPropType> = ({ index, row, fields }) => {
  return (
    <tr className="bg-white border-b last:border-b-0 hover:bg-indigo-50 transition">
      <td className="px-6 py-4 text-center font-medium text-gray-900">
        {index + 1}
      </td>
      {
        fields.map(field => (
          <td className={`px-6 py-4${isDigitField(field) ? " font-mono" : ""}`}>
            {field.type === "OPTION" ?
              <span className="px-2 py-1 text-xs font-medium text-gray-600 bg-transparent border border-gray-300 rounded-full">
                {fieldToDisplayValue(row, field)}
              </span>
              : fieldToDisplayValue(row, field)}
            {field.type === "USER" && !row[field.fieldName].value.userId && (
              <span className="ml-2 px-2 py-1 text-xs font-medium text-gray-600 bg-gray-100 rounded-full">
                Anonymous
              </span>
            )}
          </td>
        ))
      }
    </tr>

  )
}

export default LeaderBoardContent
