import { SendWarenlieferung } from "@/api/Warenlieferung";
import { Button } from "@/components/ui/button";
import { useState } from "react";

export default function Warenlieferung() {
  const [erstellt, setErstellt] = useState(false);
  const [versendet, setVersendet] = useState(false);
  const [loading, setLoading] = useState(false);

  return (
    <div className="flex flex-col gap-8">
      <h1 className="text-center">Warenlieferung</h1>

      {erstellt ? (
        <>
          <Button
            onClick={async () => {
              setLoading(true);
              await SendWarenlieferung();
              setVersendet(true);
              setLoading(false);
            }}
            disabled={loading}
          >
            {loading ? "Sendet ..." : "Senden"}
          </Button>
        </>
      ) : (
        <>
          <Button
            onClick={async () => {
              setLoading(true);
              setErstellt(true);
              setLoading(false);
            }}
            disabled={loading}
          >
            {loading ? "Erstellt ... " : "Erstellen"}
          </Button>
        </>
      )}

      {versendet && (
        <h2 className="text-center text-pretty">Mail Erfolgreich versendet</h2>
      )}
    </div>
  );
}
