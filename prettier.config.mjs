/**
 * Shared Prettier configuration for the monorepo
 */
export default {
  printWidth: 100,
  tabWidth: 2,
  useTabs: false,
  semi: true,
  singleQuote: true,
  trailingComma: 'all',
  bracketSpacing: true,
  arrowParens: 'always',
  endOfLine: 'lf',

  // JSX
  jsxSingleQuote: false,
  bracketSameLine: false,

  // Plugins
  plugins: [
    '@ianvs/prettier-plugin-sort-imports',
    'prettier-plugin-tailwindcss',
    'prettier-plugin-packagejson',
  ],

  // Import sorting
  importOrder: [
    '^react$',
    '^next(/.*)?$',
    '<THIRD_PARTY_MODULES>',
    '',
    '^@qreate/(.*)$',
    '^@/(.*)$',
    '^[./]',
  ],
  importOrderSeparation: false,
  importOrderSortSpecifiers: true,
};
