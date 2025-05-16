import { NavLink } from "react-router";
import { Button } from "./ui/button";

export default function BackButton({ href = "/" }: { href?: string }) {
  return (
    <Button variant={"secondary"} asChild className="ms-2">
      <NavLink to={href}>Zurück</NavLink>
    </Button>
  );
}
