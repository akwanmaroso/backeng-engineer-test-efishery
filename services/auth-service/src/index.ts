import express, { Application } from "express";
import cors from "cors";
import router from "./routes/index.route";
import config from "./config/config";

const app: Application = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cors());

app.use("/api/v1", router);

try {
  app.listen(config.port, () => {
    console.log(`running on port ${config.port}`);
  });
} catch (error) {
  console.error(error);
}
