import { GetCmsCounter, type GetCmsCounterResponse } from "@/api/cms";
import { useEffect, useState } from "react";

export default function Overview() {
    const [counts, setCounts] = useState<GetCmsCounterResponse | undefined>(undefined)
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")

    useEffect(() => {
        (async () => {
            setLoading(true)
            const res = await GetCmsCounter()
            if (res == null){
                setError("Keine Daten gefunden")
                setLoading(false)
                return
            }  
            setCounts(res)
            setLoading(false)
        })()
    }, [])

    return (
        <div>
            {loading ? <>Lädt...</> : (
                <>
                    {error && <>FEHLER!: {error}</>}
                    <p>Count:</p>
                    <p>Abteilungen: {counts?.Abteilungen}</p>
                    <p>Angebote: {counts?.Angebote}</p>
                    <p>Jobs: {counts?.Jobs}</p>
                    <p>Mitarbeiter: {counts?.Mitarbeiter}</p>
                    <p>Partner: {counts?.Partner}</p>
                </>
            )}
        </div>
    )
}