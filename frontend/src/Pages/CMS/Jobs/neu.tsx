import BackBtn from "../_components/back-button";
import JobForm from "./_components/form";

export default function JobAnlegen() {
  return (
    <div>
      {" "}
      <BackBtn href="/CMS/Jobs" />
      <h1 className="text-center">Neuen Job anlegen</h1>
      <JobForm />
    </div>
  );
}
