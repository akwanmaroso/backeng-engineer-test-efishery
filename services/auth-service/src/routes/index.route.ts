import express from "express";
import { registerAuthRouter } from "./user/auth.route";
import AuthService from "../services/auth.service";
import AuthController from "../controllers/auth.controller";
import UserRepository from "../repositories/user.repository";

const router = express.Router();

const userRepository = new UserRepository();
const authService = new AuthService(userRepository);
const userController = new AuthController(authService);

const authRoute = express.Router();

registerAuthRouter(authRoute, userController);
router.use("/auth", authRoute);

export default router;
