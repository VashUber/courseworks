import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve, dirname } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

export default defineConfig((cfg) => ({
  plugins: [vue()],

  build: {
    outDir: resolve(__dirname, "../static/spa_app"),
    sourcemap: cfg.mode === "dev",
    emptyOutDir: true,
  },

  base: "/spa_app/",
}));
