import BackBtn from "../_components/back-button";
import AbteilungenTable from "./_components/table";

export default function AbteilungenÜbersicht() {
  return (
    <>
      <BackBtn />
      <h1 className="text-center">CMS - Abteilungen</h1>
      <AbteilungenTable />
    </>
  );
}
