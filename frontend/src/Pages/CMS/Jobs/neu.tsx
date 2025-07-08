import { lazy } from "react";
import BackBtn from "../_components/back-button";
const JobForm = lazy(() => import("./_components/form"));

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
