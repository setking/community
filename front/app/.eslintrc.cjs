module.exports = {
  root: true,

  env: {
    browser: true,
    node: true,
    es2021: true,
  },

  parser: 'vue-eslint-parser',

  parserOptions: {
    parser: '@typescript-eslint/parser',
    ecmaVersion: 'latest',
    sourceType: 'module',
    extraFileExtensions: ['.vue'],
  },

  extends: [
    'eslint:recommended',

    // Vue 3 + <script setup>
    'plugin:vue/vue3-recommended',

    // TypeScript 基础规则（不做类型检查）
    'plugin:@typescript-eslint/recommended',

    // ⚠️ 一定放最后
    'plugin:prettier/recommended',
  ],

  rules: {
    /* ========================
       TypeScript 放宽
       ======================== */

    // Vuetify / Vue 常见写法
    '@typescript-eslint/no-unused-vars': [
      'warn',
      { argsIgnorePattern: '^_', varsIgnorePattern: '^_' },
    ],

    '@typescript-eslint/no-explicit-any': 'off',

    /* ========================
       Vuetify / template 放宽
       ======================== */

    'vue/multi-word-component-names': 'off',
    'vue/max-attributes-per-line': 'off',
    'vue/attributes-order': 'off',
    'vue/v-slot-style': 'off',
    'vue/valid-v-slot': 'off',

    /* ========================
       通用
       ======================== */

    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
  },
}
