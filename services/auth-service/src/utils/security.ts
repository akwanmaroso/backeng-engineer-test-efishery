import { Request } from "express";

export const parseTokenFromRequest = (req: Request): string => {
  const authorizationHeaders = req.headers.authorization?.split(" ");
  if (authorizationHeaders?.length !== 2) {
    throw new Error("invalid format token");
  }

  const token = authorizationHeaders[1];
  return token;
};
