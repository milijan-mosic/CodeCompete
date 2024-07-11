import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: true,
    port: 3000,
  },
  resolve: {
    alias: {
      "@services": "/src/services/index",
      "@assets": "/src/assets/index",
      "@components": "/src/components/index",
      "@pages": "/src/pages/index",
      "@store": "/src/store",
      "@types": "/src/types",
      "@utils": "/src/utils",
    },
  },
});
