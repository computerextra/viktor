import type { userdata } from "@wails/go/models";
import {
  Login as APILogin,
  Logout as APILogout,
  CheckSession as APICheck,
} from "@wails/go/main/App";

export const Login = async (
  mail: string,
  password: string
): Promise<userdata.UserData | undefined> => {
  const res = await APILogin(mail, password);
  return res;
};

export const Logout = async (): Promise<boolean> => {
  const res = await APILogout();
  return res;
};

export const CheckSession = async (): Promise<
  userdata.UserData | undefined
> => {
  const res = await APICheck();
  return res;
};
