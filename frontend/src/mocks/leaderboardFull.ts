import { LeaderboardFull } from "@/types/leaderboard";

export const MockLeaderboardFull: LeaderboardFull = {
  id: 10,
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
      name: "score",
      type: "NUMBER",
      for_rank: true,
      required: true,
      fieldOrder: 2,
    },
    {
      name: "igt",
      type: "DURATION",
      required: false,
      fieldOrder: 3,
    },
    {
      name: "version",
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
      username: "Huyen",
      verified: true,
      fields: {
        score: 100_000_000,
        igt: 131_634_934,
        version: "1.1.2",
      },
    },
    {
      id: "fdoiioasj",
      updatedAt: "11-02-2023",
      createdAt: "11-02-2024",
      username: "Test",
      verified: true,
      fields: {
        score: 99_098_124,
        igt: 234_721_796,
        version: "2.0.0",
      },
    },
    {
      id: "oiis",
      updatedAt: "11-02-2023",
      createdAt: "11-02-2024",
      username: "RandomAnon",
      verified: true,
      fields: {
        score: 20_828_227,
        igt: 357_816_126,
        version: "1.2.1",
      },
    },
  ],
};
