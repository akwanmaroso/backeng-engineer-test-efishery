import Joi from "joi";
import "dotenv/config";

const envVarSchema = Joi.object()
  .keys({
    NODE_ENV: Joi.string().valid("production", "development").required(),
    PORT: Joi.number().default(3002),
    JWT_SECRET: Joi.string().required().description("JWT secret key"),
  })
  .unknown();

const { value: envVars, error } = envVarSchema
  .prefs({
    errors: { label: "key" },
  })
  .validate(process.env);

if (error) {
  throw new Error(`Config validation error: ${error.message}`);
}

const config = {
  env: envVars.NODE_ENV,
  port: envVars.PORT,
  jwt: {
    secret: envVars.JWT_SECRET,
  },
};

export default config;
