import { CheckSession } from "@api/userdata";
import type { userdata } from "@wails/go/models";
import { useEffect, useState } from "react";

export default function useSession() {
  const [session, setSession] = useState<undefined | userdata.UserData>(
    undefined
  );

  useEffect(() => {
    async function x() {
      const res = await CheckSession();
      setSession(res);
    }
    x();
  }, []);

  return session;
}
