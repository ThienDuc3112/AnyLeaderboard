import { LeaderboardFull } from "@/types/leaderboard";

export const MockLeaderboardFull: LeaderboardFull = {
  id: "dasflkjjlkadfs",
  name: "Bloon something the game",
  description: "Something something test game leaderboard",
  coverImageUrl:
    "https://tinfoil.media/i/0100E680149DC000/0/0/749e353633f36e5e567be03698795085b13e72340949847780fefbd8450fa8bb",
  externalLinks: [
    {
      displayValue: "Community Discord",
      url: "discord.gg/lmaoBTD",
    },
  ],
  allowAnonymous: true,
  requiredVerification: false,
  entriesCount: 3,
  fields: [
    {
      name: "Player",
      fieldName: "player",
      type: "USER",
      allowAnonymous: true,
      required: true,
      fieldOrder: 1,
    },
    {
      name: "Score",
      fieldName: "score",
      type: "NUMBER",
      for_rank: true,
      required: true,
      fieldOrder: 2,
    },
    {
      name: "IGT",
      fieldName: "igt",
      type: "DURATION",
      required: false,
      fieldOrder: 3,
    },
    {
      name: "Version",
      fieldName: "version",
      type: "OPTION",
      fieldOrder: 4,
      options: [
        "1.0.0",
        "1.0.1",
        "1.0.2",
        "1.0.3",
        "1.0.4",
        "1.1.0",
        "1.1.1",
        "1.1.2",
        "1.2.0",
        "1.2.1",
        "2.0.0",
        "2.0.1",
      ],
    },
  ],
  data: [
    {
      id: "afdosiioafsdj",
      updatedAt: "11-02-2023",
      createdAt: "11-02-2024",
      fields: {
        player: {
          value: {
            username: "Huyen",
            userId: "978241308792143",
          },
        },
        score: {
          value: 100_000_000,
        },
        igt: {
          value: 131_634_934,
        },
        version: {
          value: "1.1.2",
        },
      },
    },
    {
      id: "fdoiioasj",
      updatedAt: "11-02-2023",
      createdAt: "11-02-2024",
      fields: {
        player: {
          value: {
            username: "Isab",
            userId: "835685283932598",
          },
        },
        score: {
          value: 99_098_124,
        },
        igt: {
          value: 234_721_796,
        },
        version: {
          value: "2.0.0",
        },
      },
    },
    {
      id: "oiis",
      updatedAt: "11-02-2023",
      createdAt: "11-02-2024",
      fields: {
        player: {
          value: {
            username: "RandomAnon",
          },
        },
        score: {
          value: 20_828_227,
        },
        igt: {
          value: 357_816_126,
        },
        version: {
          value: "1.2.1",
        },
      },
    },
  ],
};
