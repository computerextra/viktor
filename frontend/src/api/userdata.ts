import {
  CheckSession as APICheckSession,
  Login as APILogin,
  Logout as APILogout,
} from "@bindings/viktor/backend/app";
import type { UserData } from "@bindings/viktor/userdata/models";

export const Login = async (
  mail: string,
  password: string
): Promise<UserData | null> => {
  const res = await APILogin(mail, password);

  return res;
};

export const Logout = async (): Promise<boolean> => {
  const res = await APILogout();
  return res;
};

export const CheckSession = async (): Promise<UserData | null> => {
  const res = await APICheckSession();
  return res;
};
