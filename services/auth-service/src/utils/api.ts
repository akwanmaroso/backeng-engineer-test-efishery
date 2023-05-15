import { Response } from "express";
import httpStatusCodes from "http-status-codes";

export interface IOverrideRequest {
  code: number;
  message: string;
}
export default class ApiResponse {
  static result = (
    res: Response,
    data: object | null,
    status: number = 200
  ) => {
    res.status(status).json({ data, success: true });
  };

  static error = (
    res: Response,
    status: number = 400,
    error: string = httpStatusCodes.getStatusText(status)
  ) => {
    res.status(status).json({
      error: { message: error },
      success: false,
    });
  };
}
