import Dropdown from "@/components/ui/Dropdown";
import Input from "@/components/ui/Input";
import { Field } from "@/types/leaderboard";
import { useFormikContext } from "formik";
import React, { useEffect, useState } from "react";
import {
  formatMsToDuration,
  formatMsToTimestamp,
  parseDurationToMs,
  parseTimestampToMs,
} from "./utils";

interface PropType {
  field: Field;
}

const FieldInput: React.FC<PropType> = ({ field }) => {
  const p = useFormikContext<Record<string, any>>();
  switch (field.type) {
    case "TEXT":
      return (
        <Input
          type="text"
          placeholder={`Enter ${field.name}`}
          name={field.name}
          value={p.values[field.name]}
          onChange={p.handleChange}
          onBlur={p.handleBlur}
        />
      );

    case "NUMBER":
      return (
        <Input
          type="number"
          placeholder={`Enter ${field.name}`}
          name={field.name}
          value={p.values[field.name]}
          onChange={p.handleChange}
          onBlur={p.handleBlur}
        />
      );

    case "OPTION":
      return (
        <div className="h-8 flex flex-1">
          <Dropdown
            options={field.options.map((o) => ({ text: o, value: o }))}
            name={field.name}
            value={p.values[field.name]}
            onChange={p.handleChange}
            onBlur={p.handleBlur}
          />
        </div>
      );

    case "DURATION": {
      const value = p.values[field.name];
      const [localInput, setLocalInput] = useState(
        typeof value === "number" && !isNaN(value)
          ? formatMsToDuration(value)
          : "",
      );

      // Sync if form is reset or value changes externally
      useEffect(() => {
        if (typeof value === "number") {
          setLocalInput(formatMsToDuration(value));
        }
      }, [value]);

      const handleBlur = () => {
        const ms = parseDurationToMs(localInput);
        p.setFieldValue(field.name, ms || 0); // fallback to 0 on bad input
        p.handleBlur({ target: { name: field.name } });
      };

      return (
        <Input
          type="text"
          placeholder="HH:MM:SS.mmm"
          name={field.name}
          value={localInput}
          onChange={(e) => setLocalInput(e.target.value)}
          onBlur={handleBlur}
        />
      );
    }

    case "TIMESTAMP": {
      const value = p.values[field.name];
      const [localInput, setLocalInput] = useState(
        typeof value === "number" && !isNaN(value)
          ? formatMsToTimestamp(value)
          : "",
      );
      useEffect(() => {
        if (typeof value === "number") {
          setLocalInput(formatMsToTimestamp(value));
        }
      }, [value]);

      const handleBlur = () => {
        const ms = parseTimestampToMs(localInput);
        p.setFieldValue(field.name, ms || 0); // fallback to 0 on bad input
        p.handleBlur({ target: { name: field.name } });
      };

      return (
        <Input
          type="text"
          name={field.name}
          placeholder="YYYY-MM-DD HH:MM:SS"
          value={localInput}
          onChange={(e) => setLocalInput(e.target.value)}
          onBlur={handleBlur}
        />
      );
    }
    default:
      return null;
  }
};

export default FieldInput;
