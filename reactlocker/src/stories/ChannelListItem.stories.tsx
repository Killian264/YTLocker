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
			"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
		title: "DogeLog",
		description: "Videos showing Ben Awad as he builds dogehouse.",
		created: new Date(),
		videos: [],
	},
];

const Mocked: Story<{}> = ({ ...props }) => {
	let url = "https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK";

	return (
		<Card>
			<div></div>
			<div>
				<ChannelListItem
					url={url}
					channel={channels[0]}
					colors={["red-1", "yellow-1"]}
				></ChannelListItem>
				<ChannelListItem
					url={url}
					channel={channels[0]}
					colors={["red-1"]}
					className="mt-3"
				></ChannelListItem>
				<ChannelListItem
					url={url}
					channel={channels[0]}
					colors={["yellow-1"]}
					className="mt-3"
				></ChannelListItem>
			</div>
		</Card>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
