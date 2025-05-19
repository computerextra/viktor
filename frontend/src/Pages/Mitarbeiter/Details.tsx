import BackButton from "@/components/BackButton";
import { Button } from "@/components/ui/button";
import { Mitarbeiter } from "@api/db";
import type { db } from "@wails/go/models";
import { useEffect, useState } from "react";
import { Link, useParams } from "react-router";

export default function MitarbeiterDetails() {
  const { id } = useParams();
  const [mitarbeiter, setMitarbeiter] = useState<db.Mitarbeiter | undefined>(
    undefined
  );

  useEffect(() => {
    async function x() {
      if (id == null) return;
      const i = parseInt(id);
      const res = await Mitarbeiter.Get(i);
      setMitarbeiter(res);
    }

    x();
  }, [id]);
  return (
    <>
      <BackButton href="/Mitarbeiter" />
      <Button asChild>
        <Link to={`/Mitarbeiter/${id}/Bearbeiten`}>Bearbeiten</Link>
      </Button>
      <div>{mitarbeiter?.Name}</div>
    </>
  );
}
