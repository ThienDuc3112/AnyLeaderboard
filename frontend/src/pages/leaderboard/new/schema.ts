import * as y from "yup";

const LbMetadataSchema = y.object({
  name: y.string().required("Name is required"),
  description: y.string().max(256, "Cannot exceed 256 characters"),
  coverImageUrl: y.string().url("Must be an image url"),
  externalLinks: y
    .array(
      y.object().shape({
        displayValue: y
          .string()
          .required("Display value is required")
          .max(32, "Cannot exceed 32 characters"),
        url: y.string().url("Must be a valid url").required("url is required"),
      }),
    )
    .max(5, "Cannot exceed 5 external links")
    .required("external links must exist"),
  uniqueSubmission: y.boolean(),
  allowAnonymous: y.boolean().when("uniqueSubmission", {
    is: true,
    then: (s) =>
      s.isFalse("Cannot allow anonymous when unique submission is on"),
    otherwise: (s) => s,
  }),
  requiredVerification: y.boolean(),
});

export const SubmitSchema = LbMetadataSchema.shape({
  fields: y
    .array(
      y.object().shape({
        name: y
          .string()
          .required("Field name is required")
          .min(1, "Field name must be atleast 1 character")
          .max(32, "Cannot exceed 32 characters"),
        required: y.boolean().when("forRank", {
          is: true,
          then: (s) => s.isTrue("For rank field cannot be empty"),
        }),
        hidden: y.boolean(),
        type: y
          .string()
          .required("Field type must be specified")
          .oneOf(["TEXT", "NUMBER", "DURATION", "TIMESTAMP", "OPTION"]),
        forRank: y
          .boolean()
          .required("For rank is required")
          .when("type", {
            is: "OPTION",
            then: (schema) => schema.isFalse("Option field cannot be rank"),
            otherwise: (schema) => schema,
          })
          .when("type", {
            is: "TEXT",
            then: (s) => s.isFalse("Text field cannot be rank"),
          }),
        options: y.string().when("type", {
          is: "OPTION",
          then: (s) =>
            s
              .required("Option must exist when field type is OPTION")
              .min(1, "At least 1 option must exist"),
        }),
      }),
    )
    .required("Atleast 1 field must exist")
    .test(
      "oneForRank",
      "There must be and can only be 1 for rank field",
      (fields) => {
        const count = fields.reduce(
          (acc: number, field) => (field.forRank ? acc + 1 : acc),
          0,
        );
        return count == 1;
      },
    )
    .test(
      "onePublicField",
      "There must be atleast 1 for non hidden field",
      (fields) => {
        const count = fields.reduce(
          (acc: number, field) => (field.hidden ? acc : acc + 1),
          0,
        );
        return count >= 1;
      },
    )
    .min(1, "Atleast 1 field must exist")
    .max(10, "A leaderboard cannot have more than 10 fields"),
});

export type SubmitType = y.InferType<typeof SubmitSchema>;
