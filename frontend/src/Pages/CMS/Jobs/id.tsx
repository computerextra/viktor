import { useParams } from "react-router";
import BackBtn from "../_components/back-button";
import JobForm from "./_components/form";

export default function JobBearbeiten() {
  const { id } = useParams();

  return (
    <>
      <BackBtn href="/CMS/Jobs" />
      <h1 className="text-center">Job Bearbeiten</h1>
      <JobForm id={id} />
    </>
  );
}
