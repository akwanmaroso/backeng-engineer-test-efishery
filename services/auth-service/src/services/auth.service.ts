import { UserInput } from "../entities/user/user.entity";
import UserRepository from "../repositories/user.repository";
import { createRandomPassword } from "../utils/generate";
import TokenService from "./token.service";

export interface LoginPayload {
  phoneNumber: string;
  password: string;
}

export interface RegisterPayload {
  phoneNumber: string;
  role: string;
  name: string;
}

export default class AuthService {
  private userRepository: UserRepository;
  constructor(userRepository: UserRepository) {
    this.userRepository = userRepository;
  }

  public async login({ phoneNumber, password }: LoginPayload) {
    try {
      const user = await this.userRepository.findByPhoneNumber(phoneNumber);
      if (!user) {
        throw new Error("user not found");
      }

      if (user.password !== password) {
        throw new Error("invalid credential");
      }

      const token = TokenService.generateAuthToken(user);
      return { ...token, user };
    } catch (err) {
      console.error(err);
      throw err;
    }
  }

  public async register({ name, role, phoneNumber }: RegisterPayload) {
    try {
      const newUser: UserInput = {
        name,
        role,
        phoneNumber,
        password: createRandomPassword(4),
      };
      const user = await this.userRepository.create(newUser);
      return user;
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
}
