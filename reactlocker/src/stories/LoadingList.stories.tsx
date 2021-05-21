import { Story } from "@storybook/react";
import { LoadingList } from "../components/LoadingList";

export default {
	title: "LoadingList",
	component: LoadingList,
};

const Mocked: Story<{}> = ({ ...props }) => {
	return <LoadingList limit={5}></LoadingList>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
