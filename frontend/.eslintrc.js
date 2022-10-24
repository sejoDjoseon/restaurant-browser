module.exports = {
  root: true,
  extends: ['react-app', 'react-app/jest', 'prettier'],
  parser: '@typescript-eslint/parser',
  plugins: ['@typescript-eslint', 'import'],
  rules: {
    // note you must disable the base rule as it can report incorrect errors
    'no-shadow': 'off',
    '@typescript-eslint/no-shadow': ['error'],
    'no-duplicate-imports': 'error',
    'no-redeclare': 'error',
    'import/no-anonymous-default-export': 'off',
    'import/order': [
      'error',
      {
        groups: ['builtin', 'external', 'internal', 'index'],
        pathGroups: [
          {
            pattern: '*(react|react-native)',
            group: 'builtin',
            position: 'before',
          },
          {
            pattern:
              '*(@components|@app|@libs|@routes|@hocs|@screens|@models|@services|@assets|@locale|@theme|@utils|@providers)/**',
            group: 'internal',
          },
        ],
        pathGroupsExcludedImportTypes: ['react'],
        'newlines-between': 'always',
        alphabetize: {
          order: 'asc',
          caseInsensitive: true,
        },
      },
    ],
    'array-callback-return': 'error',
    complexity: ['error', 11], //objective 5, ideally 3
    eqeqeq: 'error',
    'max-statements': ['error', 20], //objective 7, ideally 5
    'max-statements-per-line': [
      'error',
      {
        max: 1,
      },
    ],
    'max-nested-callbacks': ['error', 5], //objective 3, ideally 2
    'max-depth': [
      'error',
      {
        max: 3, //objective 2, ideally 1
      },
    ],
    'max-lines': ['error', 500], //objective 200, ideally 100-150
    'no-eval': 'error',
    'no-return-assign': 'error',
    'no-param-reassign': 'error',
    'no-var': 'error',
    'prefer-const': 'error',
    'max-lines-per-function': ['error', 185], //Objective 40, ideally 25
    'max-params': ['error', 6], //objective5, ideally 3
  },
}
