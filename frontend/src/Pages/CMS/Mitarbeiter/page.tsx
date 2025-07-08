import BackBtn from "../_components/back-button";
import MitarbeiterTable from "./_components/table";

export default function MitarbeiterÜbersicht() {
  return (
    <>
      <BackBtn />
      <h1 className="text-center">CMS - Mitarbeiter</h1>
      <MitarbeiterTable />
    </>
  );
}
