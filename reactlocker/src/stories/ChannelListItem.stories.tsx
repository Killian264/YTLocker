import { Story } from "@storybook/react";
import { Card } from "../components/Card";
import { ChannelListItem } from "../components/ChannelListItem";
import { Channel } from "../shared/types";

export default {
	title: "ChannelListItem",
	component: ChannelListItem,
};

const channels: Channel[] = [
	{
		id: 932423423,
		youtubeId: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
		thumbnailUrl:
			"https://yt3.ggpht.com/ytc/AAUvwngwzt6aURebMaKQpNkT9WpY3A3z0rtocMufWQwLxA=s800-c-k-c0x00ffffff-no-rj-mo",
		title: "DogeLog",
		description: "Videos showing Ben Awad as he builds dogehouse.",
		created: new Date(),
		videos: [],
	},
];

const Mocked: Story<{}> = ({ ...props }) => {
	let url = "https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK";

	return (
		<div>
			<Card>
				<div></div>
				<div>
					<ChannelListItem
						url={url}
						channel={channels[0]}
						colors={["red-1", "yellow-1"]}
						mode="normal"
						remove={() => console.log("hello")}
					></ChannelListItem>
					<ChannelListItem
						url={url}
						channel={channels[0]}
						colors={["red-1"]}
						className="mt-3"
						mode="normal"
						remove={() => console.log("hello")}
					></ChannelListItem>
					<ChannelListItem
						url={url}
						channel={channels[0]}
						colors={["yellow-1"]}
						className="mt-3"
						mode="delete"
						remove={() => console.log("hello")}
					></ChannelListItem>
				</div>
			</Card>
		</div>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
