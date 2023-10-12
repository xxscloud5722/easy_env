module.exports = {
  env: {
    browser: true
  },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:react/recommended',
    // 'plugin:jsx-a11y/recommended',
    // 'plugin:react-hooks/recommended',
    'airbnb-base'
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 2022,
    sourceType: 'module',
    requireConfigFile: false,
    babelOptions: {
      presets: ['@babel/preset-react']
    }
  },
  plugins: [
    'react',
    '@typescript-eslint'
    // 'jsx-a11y',
    // 'react-hooks'
  ],
  settings: {
    react: {
      version: '18'
    }
  },
  rules: {
    'class-methods-use-this': 'off',
    'padded-blocks': 'off',
    'max-classes-per-file': ['error', 10],
    'import/extensions': 'off',
    'import/no-unresolved': 'off',
    'lines-between-class-members': 'off',
    'no-unused-vars': 'off',
    'import/prefer-default-export': 'off',
    'no-restricted-syntax': 'off',
    'no-continue': 'off',
    'max-len': ['error', 180],
    'no-param-reassign': 'off',
    'vue/no-v-model-argument': 'off',
    'arrow-body-style': 'off',
    'vue/no-multiple-template-root': 'off',
    'import/no-extraneous-dependencies': 'off',
    'no-case-declarations': 'off',
    'default-case': 'off',
    'no-return-assign': 'off',
    'indent': ['error', 2],
    'no-debugger': 'off',
    'comma-dangle': ['error', 'never'],
    'object-curly-newline': 'off',
    'no-use-before-define': 'off',
    'no-shadow': 'off',
    'vue/no-v-for-template-key': 'off',
    'linebreak-style': [0, 'error', 'windows'],
    'no-extra-semi': 2,
    'no-console': 'off',
    'spaced-comment': 'off',
    'react/display-name': 'off',
    'default-param-last': 'off',
    'prefer-template': 'off',
    'arrow-parens': 'off',
    'react/prop-types': 'off',
    '@typescript-eslint/no-unused-vars': ['warn', {
      vars: 'all',
      args: 'after-used',
      ignoreRestSiblings: true,
      varsIgnorePattern: '^_',
      argsIgnorePattern: '^_'
    }],
    '@typescript-eslint/no-explicit-any': 'off',
    'no-empty-interface': 'off',
    '@typescript-eslint/ban-types': 'off'
  }
};
