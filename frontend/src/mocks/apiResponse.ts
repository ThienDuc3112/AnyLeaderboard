import { LeaderboardFull } from "@/types/leaderboard";

export const MockApiResponse: LeaderboardFull = {
  name: "Bloon something the game",
  description: "Something something test game leaderboard",
  coverImageUrl: "https://tinfoil.media/i/0100E680149DC000/0/0/749e353633f36e5e567be03698795085b13e72340949847780fefbd8450fa8bb",
  externalLinks: [
    {
      displayValue: "Community Discord",
      url: "discord.gg/lmaoBTD",
      icon: "discord"
    }
  ],
  fields: [
    {
      name: "Player",
      fieldName: "player",
      type: "USER",
      allowAnonymous: true
    },
    {
      name: "Score",
      fieldName: "score",
      type: "INTEGER",
      defaultSort: true
    },
    {
      name: "IGT",
      fieldName: "igt",
      type: "DURATION",
    },
    {
      name: "Version",
      fieldName: "version",
      type: "OPTION",
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
      ]
    },
  ],
  data: [
    {
      player: {
        value: {
          username: "Huyen",
          userId: "978241308792143"
        }
      },
      score: {
        value: 100_000_000
      },
      igt: {
        value: 131_634_934
      },
      version: {
        value: "1.1.2"
      }
    },
    {
      player: {
        value: {
          username: "Isab",
          userId: "835685283932598"
        }
      },
      score: {
        value: 99_098_124
      },
      igt: {
        value: 234_721_796
      },
      version: {
        value: "2.0.0"
      }
    },
    {
      player: {
        value: {
          username: "RandomAnon",
        }
      },
      score: {
        value: 20_828_227
      },
      igt: {
        value: 357_816_126
      },
      version: {
        value: "1.2.1"
      }
    },
  ]
}
