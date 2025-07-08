import BackBtn from "../_components/back-button";
import AngeboteTable from "./_components/table";

export default function AngeboteÜbersicht() {
  return (
    <>
      <BackBtn />
      <h1 className="text-center">CMS - Angebote</h1>
      <AngeboteTable />
    </>
  );
}
