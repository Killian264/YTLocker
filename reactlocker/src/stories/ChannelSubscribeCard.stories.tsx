import { Story } from "@storybook/react";
import { ChannelSubscribeCard } from "../components/ChannelSubscribeCard";
import { Channel } from "../shared/types";

export default {
	title: "ChannelSubscribeCard",
	component: ChannelSubscribeCard,
};

const channel: Channel = {
	id: 932423423,
	youtubeId: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
	thumbnailUrl:
		"https://yt3.ggpht.com/ytc/AAUvwngwzt6aURebMaKQpNkT9WpY3A3z0rtocMufWQwLxA=s800-c-k-c0x00ffffff-no-rj-mo",
	title: "DogeLog",
	description: "Videos showing Ben Awad as he builds dogehouse.",
	created: new Date(),
	videos: [],
};

const Mocked: Story<{}> = () => {
	return (
		<div>
			<ChannelSubscribeCard
				channel={channel}
				SubscribeClick={() => console.log("subscribed")}
				BackClick={() => console.log("go back")}
				SearchChannel={() => console.log("search calledd")}
			></ChannelSubscribeCard>
			<ChannelSubscribeCard
				className="mt-3"
				channel={null}
				SubscribeClick={() => console.log("subscribed")}
				BackClick={() => console.log("go back")}
				SearchChannel={() => console.log("search calledd")}
			></ChannelSubscribeCard>
		</div>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
