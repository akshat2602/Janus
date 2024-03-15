import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import eslintPlugin from "vite-plugin-eslint";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 3000,
  },
  plugins: [
    react(),
    eslintPlugin({
      cache: false,
      include: ["./src/**/*.tsx", "./src/**/*.ts"],
      fix: true,
    }),
  ],
});
