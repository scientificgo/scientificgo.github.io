/* eslint-disable @typescript-eslint/no-require-imports */

const { FlatCompat } = require("@eslint/eslintrc");

const compat = new FlatCompat({
  baseDirectory: __dirname,
});

module.exports = [
  ...compat.extends("next/core-web-vitals", "next/typescript"),
  ...compat.config({ root: true }),
];
