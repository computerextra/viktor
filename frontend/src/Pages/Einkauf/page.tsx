import EinkaufKnöppe from "./_components/btns";
import EinkaufListe from "./_components/liste";

export default function Einkauf() {
  return (
    <>
      <h1 className="text-center print:hidden">Einkaufsliste</h1>
      <h1 className="hidden text-center print:block">
        An Post / Milch denken !
      </h1>
      <div className="print:hidden">
        <EinkaufKnöppe />
      </div>
      <EinkaufListe />
    </>
  );
}
