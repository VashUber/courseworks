import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve, dirname } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

export default defineConfig((cfg) => {
  const env = loadEnv(cfg.mode, resolve(__dirname, ".."), "");
  const root = env.SPA_ROOT;

  return {
    plugins: [vue()],

    build: {
      outDir: resolve(__dirname, `../static/${root}`),
      sourcemap: cfg.mode === "dev",
      emptyOutDir: true,
    },

    base: `/${root}/`,
  };
});
