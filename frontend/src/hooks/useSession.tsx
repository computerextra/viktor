import { CheckSession } from "@/api/userdata";
import type { UserData } from "bindings/viktor/userdata/models";
import { useEffect, useState } from "react";

export default function useSession() {
  const [session, setSession] = useState<null | UserData>(null);

  useEffect(() => {
    async function x() {
      const res = await CheckSession();
      setSession(res);
    }
    x();
  }, []);

  return session;
}
