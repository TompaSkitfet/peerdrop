import { Button } from "@mui/material";
import { Box } from "@mui/system";

export default function Home() {
  return (
    <Box sx={{ height: "100%", width: "100%" }}>
      <Button>Create Session</Button>
      <Button>Join Session</Button>
    </Box>
  );
}
