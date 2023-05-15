import TokenService from "../services/token.service";
import ApiResponse from "../utils/api";
import { parseTokenFromRequest } from "../utils/security";
import { NextFunction, Request, Response } from "express";
import httpStatusCodes from "http-status-codes";

const authMiddleware = (req: Request, res: Response, next: NextFunction) => {
  try {
    const token = parseTokenFromRequest(req);
    if (!token) {
      return ApiResponse.error(
        res,
        httpStatusCodes.UNAUTHORIZED,
        "token not provided"
      );
    }
    const decodedToken = TokenService.decodeToken(token);
    console.log(decodedToken);
    // @ts-ignore
    req.userId = decodedToken.sub;
    // @ts-ignore
    req.user = decodedToken.user;
  } catch (error) {
    console.log(error);
    return ApiResponse.error(res, httpStatusCodes.FORBIDDEN, "invalid token");
  }
  next();
};
export default authMiddleware;
