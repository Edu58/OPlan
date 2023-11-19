defmodule Oplan.EventTest do
  use Oplan.DataCase

  alias Oplan.Event

  describe "event_rating" do
    alias Oplan.Event.EventRating

    import Oplan.EventFixtures

    @invalid_attrs %{event_ratings: nil, rating: nil}

    test "list_event_rating/0 returns all event_rating" do
      event_rating = event_rating_fixture()
      assert Event.list_event_rating() == [event_rating]
    end

    test "get_event_rating!/1 returns the event_rating with given id" do
      event_rating = event_rating_fixture()
      assert Event.get_event_rating!(event_rating.id) == event_rating
    end

    test "create_event_rating/1 with valid data creates a event_rating" do
      valid_attrs = %{event_ratings: "some event_ratings", rating: 42}

      assert {:ok, %EventRating{} = event_rating} = Event.create_event_rating(valid_attrs)
      assert event_rating.event_ratings == "some event_ratings"
      assert event_rating.rating == 42
    end

    test "create_event_rating/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Event.create_event_rating(@invalid_attrs)
    end

    test "update_event_rating/2 with valid data updates the event_rating" do
      event_rating = event_rating_fixture()
      update_attrs = %{event_ratings: "some updated event_ratings", rating: 43}

      assert {:ok, %EventRating{} = event_rating} =
               Event.update_event_rating(event_rating, update_attrs)

      assert event_rating.event_ratings == "some updated event_ratings"
      assert event_rating.rating == 43
    end

    test "update_event_rating/2 with invalid data returns error changeset" do
      event_rating = event_rating_fixture()
      assert {:error, %Ecto.Changeset{}} = Event.update_event_rating(event_rating, @invalid_attrs)
      assert event_rating == Event.get_event_rating!(event_rating.id)
    end

    test "delete_event_rating/1 deletes the event_rating" do
      event_rating = event_rating_fixture()
      assert {:ok, %EventRating{}} = Event.delete_event_rating(event_rating)
      assert_raise Ecto.NoResultsError, fn -> Event.get_event_rating!(event_rating.id) end
    end

    test "change_event_rating/1 returns a event_rating changeset" do
      event_rating = event_rating_fixture()
      assert %Ecto.Changeset{} = Event.change_event_rating(event_rating)
    end
  end

  describe "event_comments" do
    alias Oplan.Event.EventComment

    import Oplan.EventFixtures

    @invalid_attrs %{comment: nil}

    test "list_event_comments/0 returns all event_comments" do
      event_comment = event_comment_fixture()
      assert Event.list_event_comments() == [event_comment]
    end

    test "get_event_comment!/1 returns the event_comment with given id" do
      event_comment = event_comment_fixture()
      assert Event.get_event_comment!(event_comment.id) == event_comment
    end

    test "create_event_comment/1 with valid data creates a event_comment" do
      valid_attrs = %{comment: "some comment"}

      assert {:ok, %EventComment{} = event_comment} = Event.create_event_comment(valid_attrs)
      assert event_comment.comment == "some comment"
    end

    test "create_event_comment/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Event.create_event_comment(@invalid_attrs)
    end

    test "update_event_comment/2 with valid data updates the event_comment" do
      event_comment = event_comment_fixture()
      update_attrs = %{comment: "some updated comment"}

      assert {:ok, %EventComment{} = event_comment} =
               Event.update_event_comment(event_comment, update_attrs)

      assert event_comment.comment == "some updated comment"
    end

    test "update_event_comment/2 with invalid data returns error changeset" do
      event_comment = event_comment_fixture()

      assert {:error, %Ecto.Changeset{}} =
               Event.update_event_comment(event_comment, @invalid_attrs)

      assert event_comment == Event.get_event_comment!(event_comment.id)
    end

    test "delete_event_comment/1 deletes the event_comment" do
      event_comment = event_comment_fixture()
      assert {:ok, %EventComment{}} = Event.delete_event_comment(event_comment)
      assert_raise Ecto.NoResultsError, fn -> Event.get_event_comment!(event_comment.id) end
    end

    test "change_event_comment/1 returns a event_comment changeset" do
      event_comment = event_comment_fixture()
      assert %Ecto.Changeset{} = Event.change_event_comment(event_comment)
    end
  end

  describe "event_followers" do
    alias Oplan.Event.EventFollower

    import Oplan.EventFixtures

    @invalid_attrs %{}

    test "list_event_followers/0 returns all event_followers" do
      event_follower = event_follower_fixture()
      assert Event.list_event_followers() == [event_follower]
    end

    test "get_event_follower!/1 returns the event_follower with given id" do
      event_follower = event_follower_fixture()
      assert Event.get_event_follower!(event_follower.id) == event_follower
    end

    test "create_event_follower/1 with valid data creates a event_follower" do
      valid_attrs = %{}

      assert {:ok, %EventFollower{} = event_follower} = Event.create_event_follower(valid_attrs)
    end

    test "create_event_follower/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Event.create_event_follower(@invalid_attrs)
    end

    test "update_event_follower/2 with valid data updates the event_follower" do
      event_follower = event_follower_fixture()
      update_attrs = %{}

      assert {:ok, %EventFollower{} = event_follower} =
               Event.update_event_follower(event_follower, update_attrs)
    end

    test "update_event_follower/2 with invalid data returns error changeset" do
      event_follower = event_follower_fixture()

      assert {:error, %Ecto.Changeset{}} =
               Event.update_event_follower(event_follower, @invalid_attrs)

      assert event_follower == Event.get_event_follower!(event_follower.id)
    end

    test "delete_event_follower/1 deletes the event_follower" do
      event_follower = event_follower_fixture()
      assert {:ok, %EventFollower{}} = Event.delete_event_follower(event_follower)
      assert_raise Ecto.NoResultsError, fn -> Event.get_event_follower!(event_follower.id) end
    end

    test "change_event_follower/1 returns a event_follower changeset" do
      event_follower = event_follower_fixture()
      assert %Ecto.Changeset{} = Event.change_event_follower(event_follower)
    end
  end
end
