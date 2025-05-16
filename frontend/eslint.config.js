import js from "@eslint/js";
import globals from "globals";
import reactHooks from "eslint-plugin-react-hooks";
import reactRefresh from "eslint-plugin-react-refresh";
import tseslint from "typescript-eslint";
import erasableSyntaxOnly from "eslint-plugin-erasable-syntax-only";

export default tseslint.config(
  { ignores: ["dist", "wailsjs"] },
  {
    extends: [erasableSyntaxOnly.configs.recommended],
    ignores: ["wailsjs"],
  },
  {
    extends: [js.configs.recommended, ...tseslint.configs.recommended],
    files: ["**/*.{ts,tsx}"],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
    },
    plugins: {
      "react-hooks": reactHooks,
      "react-refresh": reactRefresh,
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      "react-refresh/only-export-components": [
        "warn",
        { allowConstantExport: true },
      ],
    },
    ignores: ["wailsjs"],
    settings: {
      "import/resolver": {
        alias: {
          extensions: [".ts", ".js", ".tsx", ".jsx"],
          map: [
            ["@/", "./src/"],
            ["@wails/", "./wailsjs/"],
            ["@api/", "./api/"],
          ],
        },
      },
    },
  }
);
