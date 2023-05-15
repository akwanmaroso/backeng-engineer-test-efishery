import express, { Router } from "express";
import AuthController from "../../controllers/auth.controller";
import authMiddleware from "../../middlewares/authentication";

export const registerAuthRouter = (
  router: Router,
  controller: AuthController
) => {
  router.post("/login", controller.login);
  router.post("/register", controller.register);
  router.get("/current", authMiddleware, controller.current);
};
