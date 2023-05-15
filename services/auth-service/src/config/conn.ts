import { Sequelize } from "sequelize";
import config from "./config";

const isDev = config.env === "development";

const sequelizeConn = new Sequelize({
  dialect: "sqlite",
  storage: "../db.sqlite",
  logging: isDev,
});

export default sequelizeConn;
