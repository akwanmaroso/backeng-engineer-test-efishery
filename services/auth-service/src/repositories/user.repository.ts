import User, { UserInput, UserOutput } from "../entities/user/user.entity";

export default class UserRepository {
  public async create(payload: UserInput): Promise<UserOutput> {
    const user = await User.create(payload);
    return user;
  }

  public async findByPhoneNumber(phoneNumber: string): Promise<User | null> {
    const user = await User.findOne({ where: { phoneNumber } });

    return user;
  }
}
