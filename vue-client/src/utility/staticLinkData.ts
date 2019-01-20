function makeItem(
  link: string,
  urlText: string
): { link: string; urlText: string } {
  return {
    link,
    urlText
  };
}

export default function GetData() {
  return [
    makeItem("/about", "About"),
    makeItem("/projects/arch-reborn-web", "Projects//Arch Reborn [On Hold]"),
    makeItem(
      "/projects/tensorflow-linear-regression",
      "Projects//tfjs Linear Regression [On Hold]"
    ),
    makeItem(
      "/projects/herding-cats",
      "Projects//herding-cats [On-going effort]"
    ),
    makeItem(
      "/projects/websocket-01",
      "Projects//WebSocket experiment 01",
    ),
  ];
}
