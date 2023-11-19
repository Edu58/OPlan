defmodule Oplan.ProfilesTest do
  use Oplan.DataCase

  alias Oplan.Profiles

  describe "profiles" do
    alias Oplan.Profiles.Profile

    import Oplan.ProfilesFixtures

    @invalid_attrs %{
      avatar: nil,
      city: nil,
      date_of_birth: nil,
      nationality: nil,
      phone_number: nil
    }

    test "list_profiles/0 returns all profiles" do
      profile = profile_fixture()
      assert Profiles.list_profiles() == [profile]
    end

    test "get_profile!/1 returns the profile with given id" do
      profile = profile_fixture()
      assert Profiles.get_profile!(profile.id) == profile
    end

    test "create_profile/1 with valid data creates a profile" do
      valid_attrs = %{
        avatar: "some avatar",
        city: "some city",
        date_of_birth: "some date_of_birth",
        nationality: "some nationality",
        phone_number: "some phone_number"
      }

      assert {:ok, %Profile{} = profile} = Profiles.create_profile(valid_attrs)
      assert profile.avatar == "some avatar"
      assert profile.city == "some city"
      assert profile.date_of_birth == "some date_of_birth"
      assert profile.nationality == "some nationality"
      assert profile.phone_number == "some phone_number"
    end

    test "create_profile/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Profiles.create_profile(@invalid_attrs)
    end

    test "update_profile/2 with valid data updates the profile" do
      profile = profile_fixture()

      update_attrs = %{
        avatar: "some updated avatar",
        city: "some updated city",
        date_of_birth: "some updated date_of_birth",
        nationality: "some updated nationality",
        phone_number: "some updated phone_number"
      }

      assert {:ok, %Profile{} = profile} = Profiles.update_profile(profile, update_attrs)
      assert profile.avatar == "some updated avatar"
      assert profile.city == "some updated city"
      assert profile.date_of_birth == "some updated date_of_birth"
      assert profile.nationality == "some updated nationality"
      assert profile.phone_number == "some updated phone_number"
    end

    test "update_profile/2 with invalid data returns error changeset" do
      profile = profile_fixture()
      assert {:error, %Ecto.Changeset{}} = Profiles.update_profile(profile, @invalid_attrs)
      assert profile == Profiles.get_profile!(profile.id)
    end

    test "delete_profile/1 deletes the profile" do
      profile = profile_fixture()
      assert {:ok, %Profile{}} = Profiles.delete_profile(profile)
      assert_raise Ecto.NoResultsError, fn -> Profiles.get_profile!(profile.id) end
    end

    test "change_profile/1 returns a profile changeset" do
      profile = profile_fixture()
      assert %Ecto.Changeset{} = Profiles.change_profile(profile)
    end
  end
end
