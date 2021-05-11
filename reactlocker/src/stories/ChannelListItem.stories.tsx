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
		youtube: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
		thumbnail:
			"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
		title: "DogeLog",
		description: "Videos showing Ben Awad as he builds dogehouse.",
		url:
			"https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK",
	},
];

const Mocked: Story<{}> = ({ ...props }) => {
	return (
		<Card>
			<div></div>
			<div>
				<ChannelListItem channel={channels[0]}></ChannelListItem>
				<ChannelListItem
					channel={channels[0]}
					className="mt-3"
				></ChannelListItem>
				<ChannelListItem
					channel={channels[0]}
					className="mt-3"
				></ChannelListItem>
			</div>
		</Card>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
