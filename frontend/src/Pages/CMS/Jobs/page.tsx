import BackBtn from "../_components/back-button";
import JobTable from "./_components/table";

export default function JobÜbersicht() {
  return (
    <>
      <BackBtn />
      <h1 className="text-center">CMS - Jobs</h1>
      <JobTable />
    </>
  );
}
