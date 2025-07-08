import { lazy } from "react";

const GeburtstagsListe = lazy(() => import("@/components/geburtstags-liste"));

export default function Home() {
  return (
    <div>
      <h1 className="mt-5 text-center">
        Achtung: Die Seite hat noch nicht alle Funktionen.
      </h1>
      <h2 className="text-center text-3xl">
        Hier wird noch fleißig gebastelt.
      </h2>
      <section className="my-8">
        <h2 className="text-center text-3xl">Geburtstage</h2>
        <GeburtstagsListe />
      </section>
    </div>
  );
}
