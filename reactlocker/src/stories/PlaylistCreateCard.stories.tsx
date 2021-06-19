import { Story } from "@storybook/react";
import { PlaylistCreateCard } from "../components/PlaylistCreateCard";

export default {
	title: "PlaylistCreateCard",
	component: PlaylistCreateCard,
};

const Mocked: Story<{}> = () => {
	return (
		<div>
			<PlaylistCreateCard
				CreateClick={() => console.log("created")}
				BackClick={() => console.log("go back")}
			></PlaylistCreateCard>
			<PlaylistCreateCard
				className="mt-3"
				CreateClick={() => console.log("created")}
				BackClick={() => console.log("go back")}
			></PlaylistCreateCard>
		</div>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
