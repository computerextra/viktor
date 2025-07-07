import { Button } from "@/components/ui/button";
import { Link } from "react-router";

export default function BackBtn({ href }: { href?: string }) {
  return (
    <Button variant={"default"} asChild>
      <Link to={href ?? "/CMS"}>Zurück</Link>
    </Button>
  );
}
