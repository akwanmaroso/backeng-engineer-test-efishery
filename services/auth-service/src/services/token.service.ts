import moment, { Moment } from "moment";
import User from "../entities/user/user.entity";
import jwt from "jsonwebtoken";
import config from "../config/config";

export interface UserClaims {
  sub: string;
  iat: number;
  exp: number;
  user: {
    id: number;
    name: string;
    phoneNumber: string;
    role: string;
    createdAt: Date;
    updatedAt: Date;
  };
}

export default class TokenService {
  private static generateToken = (
    { id, name, phoneNumber, role, createdAt, updatedAt }: User,
    expiresIn: Moment
  ): string => {
    const payload: UserClaims = {
      sub: String(id),
      iat: moment().unix(),
      exp: expiresIn.unix(),
      user: { id, name, phoneNumber, role, createdAt, updatedAt },
    };
    return jwt.sign(payload, config.jwt.secret);
  };

  public static generateAuthToken = (user: User) => {
    const accessTokenExpires = moment().add(180, "minutes");
    const accessToken = this.generateToken(user, accessTokenExpires);

    return {
      token: accessToken,
      expires: accessTokenExpires.toDate(),
    };
  };

  public static decodeToken = (token: string): UserClaims => {
    return jwt.verify(token, config.jwt.secret) as UserClaims;
  };
}
