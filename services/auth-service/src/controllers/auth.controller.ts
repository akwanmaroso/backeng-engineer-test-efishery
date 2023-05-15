import { Response, Request } from "express";
import IController from "../types/IController";
import ApiResponse from "../utils/api";
import httpStatusCodes from "http-status-codes";
import AuthService from "../services/auth.service";

export default class AuthController {
  private authService: AuthService;

  constructor(authService: AuthService) {
    this.authService = authService;
  }

  public login: IController = async (req, res) => {
    const { phoneNumber, password } = req.body;
    try {
      const user = await this.authService.login({ phoneNumber, password });
      ApiResponse.result(res, user, httpStatusCodes.OK);
      return;
    } catch (error) {
      console.log(error);
      ApiResponse.error(res, httpStatusCodes.INTERNAL_SERVER_ERROR);
    }
  };

  public register: IController = async (req, res) => {
    const { phoneNumber, role, name } = req.body;
    try {
      const user = await this.authService.register({
        phoneNumber,
        role,
        name,
      });

      ApiResponse.result(res, user, httpStatusCodes.CREATED);
    } catch (error) {
      ApiResponse.error(res, httpStatusCodes.BAD_REQUEST);
    }
  };

  public async current(req: Request, res: Response) {
    // @ts-ignore
    const user = req.user;

    ApiResponse.result(res, user, httpStatusCodes.OK);
  }
}
