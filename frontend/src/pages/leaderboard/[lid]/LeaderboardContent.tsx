import { Entry, Field, LeaderboardFull } from "@/types/leaderboard";
import { fieldToDisplayValue } from "@/utils/fieldToDisplayValue";
import { isDigitField } from "@/utils/isDigitField";
import React from "react";
import { useNavigate } from "react-router";

interface PropType {
  data: LeaderboardFull;
}

const LeaderboardContent: React.FC<PropType> = ({ data }) => {
  return (
    <div className="px-6 pb-6">
      <div className="overflow-x-auto rounded-lg">
        <table className="w-full text-sm text-left">
          <TableHeader fields={data.fields.filter((field) => !field.hidden)} />
          <tbody>
            {data.data.map((row, index) => (
              <TableRow
                key={index}
                row={row}
                index={index}
                fields={data.fields.filter((field) => !field.hidden)}
              />
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

interface TableHeaderPropType {
  fields: Field[];
}

const TableHeader: React.FC<TableHeaderPropType> = ({ fields }) => {
  return (
    <thead className="text-xs text-gray-200 uppercase bg-indigo-600">
      <tr>
        <th scope="col" className="px-6 py-3 text-center w-12">
          #
        </th>
        <th scope="col" className="px-6 py-3">
          User
        </th>
        {fields.map((field) => (
          <th key={field.name} scope="col" className="px-6 py-3">
            {field.name}
          </th>
        ))}
      </tr>
    </thead>
  );
};

interface TableRowPropType {
  row: Entry;
  index: number;
  fields: Field[];
}

const TableRow: React.FC<TableRowPropType> = ({ index, row, fields }) => {
  const navigate = useNavigate();
  return (
    <tr
      onClick={() => {
        navigate(`entry/${row.id}`);
      }}
      className="bg-white border-b last:border-b-0 hover:bg-indigo-50 transition"
    >
      <td className="px-6 py-4 text-center font-medium text-gray-900">
        {index + 1}
      </td>
      <td className="px-6 py-4 font-medium text-gray-900">{row.username}</td>
      {fields.map((field, i) => (
        <FieldToCol key={i} field={field} row={row} />
      ))}
    </tr>
  );
};

const FieldToCol = ({ field, row }: { field: Field; row: Entry }) => {
  const value = row.fields[field.name];
  if (!value) return <td className="px-6 py-4" />;

  if (isDigitField(field))
    return (
      <td className="px-6 py-4 font-mono">{fieldToDisplayValue(row, field)}</td>
    );

  return (
    <td className={"px-6 py-4"}>
      {field.type === "OPTION" ? (
        <span className="px-2 py-1 text-xs font-medium text-gray-600 bg-transparent border border-gray-300 rounded-full">
          {fieldToDisplayValue(row, field)}
        </span>
      ) : (
        fieldToDisplayValue(row, field)
      )}
    </td>
  );
};

export default LeaderboardContent;
